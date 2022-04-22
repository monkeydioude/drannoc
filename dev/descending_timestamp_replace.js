// Usage: node descending_timestamp_replace.js <json tpl file> <delay between replace>
// Example:
// => node descending_timestamp_replace.js test.tpl.json 60000
// => will print on stdout a json with each occurence of {{timestamp}} replaced
//    by a timestamp in milliseconds starting "now", and will decrease each timestamp value
//    by 60000 ms (60s)

const fs = require("fs");
const { exit } = require("process");

const descending_timestamp_replace = (
  filePath,
  millisecondDecreasePerIteration
) => {
  try {
    const data = fs.readFileSync(`/data/${filePath}`, { encoding: "utf-8" });
    const entries = JSON.parse(data);

    if (!entries) {
      console.error("[EROR] Could not parse json in file.");
      return;
    }

    const modelDate = +new Date();
    let currDate = modelDate;

    for (const e in entries) {
      entries[e].created_at = currDate;
      currDate -= millisecondDecreasePerIteration;
    }
    console.log(JSON.stringify(entries));
  } catch (err) {
    if (err) {
      console.error(`[EROR] ${err}`);
      return;
    }
  }
};

if (require.main === module) {
  const fs = require("fs");
  const { exit } = require("process");

  if (process.argv.length < 4) {
    console.error(
      "[EROR] Target file path and timestamp delay must be given as parameters."
    );
    exit(-1);
  }
  const filePath = process.argv[2];
  const millisecondDecreasePerIteration = process.argv[3];
  descending_timestamp_replace(filePath, millisecondDecreasePerIteration);
} else {
  module.exports = descending_timestamp_replace;
}
