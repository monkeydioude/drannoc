const fs = require("fs");
const { exit } = require("process");

if (process.argv.length < 3) {
  console.error("[EROR] Target file path must be given as first parameter.");
  exit(-1);
}
const filePath = process.argv[2];

fs.readFile(`/data/${filePath}`, "utf-8", (err, data) => {
  if (err) {
    console.error(`[EROR] ${err}`);
    return;
  }
  const entries = JSON.parse(data);

  if (!entries) {
    console.error("[EROR] Could not parse json in file.");
    return;
  }

  const modelDate = +new Date();
  let currDate = modelDate;
  const millisecondDecreasePerIteration = 60000;

  for (const e in entries) {
    entries[e].created_at = currDate;
    currDate -= millisecondDecreasePerIteration;
  }
  console.log(JSON.stringify(entries));
});
