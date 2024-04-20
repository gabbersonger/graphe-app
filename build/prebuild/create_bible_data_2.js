const fs = require("fs");
const crypto = require("crypto");

const HASH_FOLDER = "/bibleVersionData";
const HASH_FILE = "/data.lock";

function hashFile(filename) {
  var f = fs.readFileSync(__dirname + filename);
  var md5 = crypto.createHash("md5");
  md5.update(f, "utf-8");
  return md5.digest("hex");
}

function checkHashFile() {
  let hashes = [];
  fs.readdirSync(__dirname + HASH_FOLDER).forEach((file) => {
    hashes.append(hashFile(HASH_FOLDER + "/" + file));
  });
  const combined_hash = hashes.join("");
  if (combined_hash == fs.readFileSync(__dirname + HASH_FILE)) return false;
  fs.writeFileSync(__dirname + HASH_FILE, combined_hash);
  return true;
}

console.log(checkHashFile());
