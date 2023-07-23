service Connectord {
  rpc ListConnectors(ListConnectorsInput) returns (ListConnectorsOutput) {}
  rpc CreateTransfer(CreateTransferInput) returns (CreateTransferOutput) {}
  rpc CancelTransfer(CancelTransferInput) returns (CancelTransferOutput) {}
  rpc ConfirmTransfer(ConfirmTransferInput) returns (ConfirmTransferOutput) {}
}

message ListConnectorsInput {
  string name = 1;
}

message ListConnectorsOutput {
  string message = 1;
}

message CreateTransferInput {
  string name = 1;
}

message CreateTransferOutput {
  string message = 1;
}

message CancelTransferInput {
  string name = 1;
}

message CancelTransferOutput {
  string message = 1;
}

message ConfirmTransferInput {
  string name = 1;
}

message ConfirmTransferOutput {
  string message = 1;
}
