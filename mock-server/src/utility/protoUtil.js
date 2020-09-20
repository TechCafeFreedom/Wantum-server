const grpc = require("grpc");
const protoLoader = require("@grpc/proto-loader");
const constants = require("./constants");

exports.getProtoDescriptor = function (protoFileName) {
  const PROTO_PATH_PREFIX = __dirname + "/../../../Wantum-ProtocolBuffer/";
  
  var packageDefinition = protoLoader.loadSync(
    PROTO_PATH_PREFIX + protoFileName,
    constants.PROTO_OPTION
  );
  return grpc.loadPackageDefinition(packageDefinition);
};