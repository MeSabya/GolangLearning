// shippy-service-consignment/proto/consignment/consignment.proto
syntax = "proto3";

package consignment; 

service ShippingService {
  rpc CreateConsignment(Consignment) returns (Response) {}
}

message Consignment {
  string id = 1;
  string description = 2;
  int32 weight = 3;
  repeated Container containers = 4;
  string vessel_id = 5;
}

message Container {
  string id = 1;
  string customer_id = 2;
  string origin = 3;
  string user_id = 4;
}

message Response {
  bool created = 1;
  Consignment consignment = 2;
}