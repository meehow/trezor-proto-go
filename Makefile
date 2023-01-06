PB = $(notdir $(wildcard trezor-common/protob/*.proto))
GO_OPT = $(addprefix --go_opt=M, $(addsuffix =./trproto, $(PB)))

all:
	protoc --go_out=. --proto_path=trezor-common/protob/ $(GO_OPT) trezor-common/protob/*.proto
