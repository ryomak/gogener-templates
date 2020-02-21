/**
 * @fileoverview gRPC-Web generated client stub for rpc
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var user_pb = require('./user_pb.js')
const proto = {};
proto.rpc = require('./room_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rpc.RoomHandlerClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rpc.RoomHandlerPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'binary';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.rpc.RoomRequest,
 *   !proto.rpc.RoomResponse>}
 */
const methodDescriptor_RoomHandler_GetRoom = new grpc.web.MethodDescriptor(
  '/rpc.RoomHandler/GetRoom',
  grpc.web.MethodType.UNARY,
  proto.rpc.RoomRequest,
  proto.rpc.RoomResponse,
  /** @param {!proto.rpc.RoomRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.RoomResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rpc.RoomRequest,
 *   !proto.rpc.RoomResponse>}
 */
const methodInfo_RoomHandler_GetRoom = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rpc.RoomResponse,
  /** @param {!proto.rpc.RoomRequest} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.RoomResponse.deserializeBinary
);


/**
 * @param {!proto.rpc.RoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rpc.RoomResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rpc.RoomResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rpc.RoomHandlerClient.prototype.getRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rpc.RoomHandler/GetRoom',
      request,
      metadata || {},
      methodDescriptor_RoomHandler_GetRoom,
      callback);
};


/**
 * @param {!proto.rpc.RoomRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rpc.RoomResponse>}
 *     A native promise that resolves to the response
 */
proto.rpc.RoomHandlerPromiseClient.prototype.getRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rpc.RoomHandler/GetRoom',
      request,
      metadata || {},
      methodDescriptor_RoomHandler_GetRoom);
};


module.exports = proto.rpc;

