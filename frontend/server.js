const express = require("express");
const path = require("path");
const app = express();
const port = 3080;
const rootDir = "/dist";

app.use(express.static(path.join(__dirname, `${rootDir}`)));

app.get(/.*/, (req, res) => {
  res.sendFile(path.join(__dirname, `${rootDir}/`));
});
app.listen(port, () => {
  console.log(`Server listening on the port::${port}`);
});
