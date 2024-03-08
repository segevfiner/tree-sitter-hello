const path = require("path");

module.exports = require("node-gyp-build")(path.join(__dirname, "../.."));

try {
  module.exports.nodeTypeInfo = require("../../src/node-types.json");
} catch (_) {}
