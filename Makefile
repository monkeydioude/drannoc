.SILENT: import generate_replaced_timestamp_tpl
latest_entries_tmp_tpl_file_name_file := /tmp/latest_entries_export
price_history_tmp_tpl_file_name_file := /tmp/price_history_export

# PARAMS (a=b):
# 	limit (integer): amount of retrieved documents.
# 	output_dir (string): export file will be generated in this directory. Gets overwritten by "output_file". Conveniant for multiple export files not overwritting each other.
# 	output_file (string): name of the export file. Overwrite "output_dir". Converniant for having a single export file, overwritting it every time.
latest_entries_export:
	$(eval tt := $(shell date +%s%N))
	$(eval tmp_tpl_file_dir := $(shell if [ -z $(output_dir) ]; then echo /tmp; else echo $(output_dir); fi))
	$(eval tmp_tpl_file := $(tmp_tpl_file_dir)/latest_entries-$(tt).json)
	$(eval limit := $(shell if [ -z $(limit) ]; then echo 10; else echo $(limit); fi))
	$(eval output_file := $(shell if [ -z $(output_file) ]; then echo $(tmp_tpl_file); else echo $(output_file); fi ))
	docker-compose exec mongo mongoexport -u $(MONGO_ADMIN_USER) -p $(MONGO_ADMIN_PWD) --authenticationDatabase=admin -d coins -c latest_entries --quiet --jsonArray --limit=$(limit) --sort="{updated_at: -1}" > $(output_file)
	echo $(output_file) > $(latest_entries_tmp_tpl_file_name_file)

latest_entries_tpl_update: latest_entries_export
	$(eval tmp_tpl_file := $(shell cat $(latest_entries_tmp_tpl_file_name_file)))
	cat $(tmp_tpl_file) | sed 's/"updated_at":[0-9]*/"updated_at":"{{timestamp}}"/gi' > dev/latest_entries.tpl.json

# PARAMS (a=b):
# 	limit (integer): amount of retrieved documents.
# 	output_dir (string): export file will be generated in this directory. Gets overwritten by "output_file". Conveniant for multiple export files not overwritting each other.
# 	output_file (string): name of the export file. Overwrite "output_dir". Converniant for having a single export file, overwritting it every time.
price_history_export:
	$(eval tt := $(shell date +%s%N))
	$(eval tmp_tpl_file_dir := $(shell if [ -z $(output_dir) ]; then echo /tmp; else echo $(output_dir); fi))
	$(eval tmp_tpl_file := $(tmp_tpl_file_dir)/price_history-$(tt).json)
	$(eval limit := $(shell if [ -z $(limit) ]; then echo 10; else echo $(limit); fi))
	$(eval output_file := $(shell if [ -z $(output_file) ]; then echo $(tmp_tpl_file); else echo $(output_file); fi ))
	docker-compose exec mongo mongoexport -u $(MONGO_ADMIN_USER) -p $(MONGO_ADMIN_PWD) --authenticationDatabase=admin -d coins -c price_history --quiet --jsonArray --limit=$(limit) --sort="{created_at: -1}" > $(output_file)
	echo $(output_file) > $(price_history_tmp_tpl_file_name_file)

price_history_tpl_update: price_history_export
	$(eval tmp_tpl_file := $(shell cat $(price_history_tmp_tpl_file_name_file)))
	cat $(tmp_tpl_file) | sed 's/"created_at":[0-9]*/"created_at":"{{timestamp}}"/gi' > dev/price_history.tpl.json

# This command is silent and won't produce a log but only one, with the resulting json after token replacement.
# PARAMS (a=b):
# 	file (string): path of the file that the {{timestamp}} token replacement will be used on.
# 	delay (integer): (optional, default 60000) amount of milliseconds the timestamp value of each {{timestamp}} will be reduced after replace.
generate_replaced_timestamp_tpl:
	$(eval pwd := $(shell pwd))
	$(eval timestamp_reduced := $(shell if [ -z $(delay) ]; then echo 60000; else echo $(delay); fi))
	@docker run --volume=$(pwd)/dev:/data node:current-alpine node /data/descending_timestamp_replace.js $(file) $(delay)

# ENV VARS
#	MONGO_ADMIN_USER (string): admin username. Should be passed as env var, not as parameters.
#	MONGO_ADMIN_USER (string): admin username. Should be passed as env var, not as parameters.
# PARAMS (a=b):
# 	file (string): path of the file to import.
#	db (string): name of the targeted Mongo DB.
#	col (string): name of the targeted Mongo collection.
#	mode (string): Optional. Default is "insert". Possible options are "insert|upsert|merge|delete"
import:
	$(eval mode := $(shell if [ -z $(mode) ]; then echo "insert"; else echo $(mode); fi ))
	$(eval container := $(shell docker-compose ps -q mongo))
	echo "Executing mongoimport using file='$(file)' db='$(db)' collection='$(col)' mode='$(mode)'"
	docker cp $(file) $(container):/tmp/import.json
	@docker-compose exec mongo mongoimport \
		-u $(MONGO_ADMIN_USER) \
		-p $(MONGO_ADMIN_PWD) \
		--jsonArray \
		-d $(db) \
		-c $(col) \
		--authenticationDatabase admin \
		--mode $(mode) \
		--file /tmp/import.json

docker_start:
	docker-compose up -d

mongo_shell: SHELL:=/bin/bash
mongo_shell:
	@source .env && docker-compose exec mongo mongo -u $(MONGO_ADMIN_USER) -p $(MONGO_ADMIN_PWD) 

dev: SHELL:=/bin/bash
dev: docker_start
	cd ./cmd/drannoc && source .env && go run main.go
