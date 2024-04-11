const fs = require("fs");

export const createBuildInfo = () => {
  const info = JSON.parse(
    fs.readFileSync(__dirname + "/../../wails.json"),
  ).info;
  fs.writeFileSync(__dirname + "/../info/title.txt", info.productName);
  fs.writeFileSync(__dirname + "/../info/version.txt", info.productVersion);
  fs.writeFileSync(__dirname + "/../info/comment.txt", info.comments);
  fs.writeFileSync(__dirname + "/../info/copyright.txt", info.copyright);
};
