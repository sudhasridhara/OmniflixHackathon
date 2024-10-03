package ems

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "ems/api/ems/ems"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "GetEvent",
					Use:            "get-event [id]",
					Short:          "Query get-event",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListEvent",
					Use:            "list-event",
					Short:          "Query list-event",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				{
					RpcMethod:      "GetTicket",
					Use:            "get-ticket [id]",
					Short:          "Query get-ticket",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},

				{
					RpcMethod:      "ListTicket",
					Use:            "list-ticket",
					Short:          "Query list-ticket",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateEvent",
					Use:            "create-event [event-name] [event-description] [start-date] [end-date] [location] [num-ft-tickets] [organizer] [ticket-price] [tickets-left]",
					Short:          "Send a create-event tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "eventName"}, {ProtoField: "eventDescription"}, {ProtoField: "startDate"}, {ProtoField: "endDate"}, {ProtoField: "location"}, {ProtoField: "numFtTickets"}, {ProtoField: "organizer"}, {ProtoField: "ticketPrice"}, {ProtoField: "ticketsLeft"}},
				},
				{
					RpcMethod:      "UpdateEvent",
					Use:            "update-event [event-name] [event-description] [start-date] [end-date] [location] [num-ft-tickets] [organizer] [ticket-price] [tickets-left] [id]",
					Short:          "Send a update-event tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "eventName"}, {ProtoField: "eventDescription"}, {ProtoField: "startDate"}, {ProtoField: "endDate"}, {ProtoField: "location"}, {ProtoField: "numFtTickets"}, {ProtoField: "organizer"}, {ProtoField: "ticketPrice"}, {ProtoField: "ticketsLeft"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "DeleteEvent",
					Use:            "delete-event [id]",
					Short:          "Send a delete-event tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "PurchaseTicket",
					Use:            "purchase-ticket [amount] [event-id] [purchaser]",
					Short:          "Send a purchase_ticket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "eventId"}, {ProtoField: "purchaser"}},
				},
				{
					RpcMethod:      "ResaleTicket",
					Use:            "resale-ticket [amount] [event-id] [purchaser] [lender] [ticket-id]",
					Short:          "Send a resale_ticket tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}, {ProtoField: "eventId"}, {ProtoField: "purchaser"}, {ProtoField: "lender"}, {ProtoField: "ticketId"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
