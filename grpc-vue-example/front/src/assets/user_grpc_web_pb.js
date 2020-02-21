/**
 * @fileoverview gRPC-Web generated client stub for rpc
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.rpc = require('./user_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.rpc.UserHandlerClient =
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
proto.rpc.UserHandlerPromiseClient =
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
 *   !proto.rpc.InvitedUser,
 *   !proto.rpc.Empty>}
 */
const methodDescriptor_UserHandler_Join = new grpc.web.MethodDescriptor(
  '/rpc.UserHandler/Join',
  grpc.web.MethodType.UNARY,
  proto.rpc.InvitedUser,
  proto.rpc.Empty,
  /** @param {!proto.rpc.InvitedUser} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.rpc.InvitedUser,
 *   !proto.rpc.Empty>}
 */
const methodInfo_UserHandler_Join = new grpc.web.AbstractClientBase.MethodInfo(
  proto.rpc.Empty,
  /** @param {!proto.rpc.InvitedUser} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.rpc.Empty.deserializeBinary
);


/**
 * @param {!proto.rpc.InvitedUser} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.rpc.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.rpc.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.rpc.UserHandlerClient.prototype.join =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/rpc.UserHandler/Join',
      request,
      metadata || {},
      methodDescriptor_UserHandler_Join,
      callback);
};


/**
 * @param {!proto.rpc.InvitedUser} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.rpc.Empty>}
 *     A native promise that resolves to the response
 */
proto.rpc.UserHandlerPromiseClient.prototype.join =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/rpc.UserHandler/Join',
      request,
      metadata || {},
      methodDescriptor_UserHandler_Join);
};


module.exports = proto.rpc;

