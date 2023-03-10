/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "escrow.escrow";

export interface MsgNewEscrow {
  creator: string;
  strategy: string;
  instigator: string;
  instigatorWager: string;
  rider: string;
  riderWager: string;
}

export interface MsgNewEscrowResponse {
  escrowIndex: string;
}

export interface MsgApprove {
  creator: string;
  escrowIndex: string;
}

export interface MsgApproveResponse {}

export interface MsgDeposit {
  creator: string;
  escrowIndex: string;
}

export interface MsgDepositResponse {}

export interface MsgConfirm {
  creator: string;
  escrowIndex: string;
  winner: string;
}

export interface MsgConfirmResponse {}

export interface MsgWithdraw {
  creator: string;
  escrowIndex: string;
}

export interface MsgWithdrawResponse {}

const baseMsgNewEscrow: object = {
  creator: "",
  strategy: "",
  instigator: "",
  instigatorWager: "",
  rider: "",
  riderWager: "",
};

export const MsgNewEscrow = {
  encode(message: MsgNewEscrow, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.strategy !== "") {
      writer.uint32(18).string(message.strategy);
    }
    if (message.instigator !== "") {
      writer.uint32(26).string(message.instigator);
    }
    if (message.instigatorWager !== "") {
      writer.uint32(34).string(message.instigatorWager);
    }
    if (message.rider !== "") {
      writer.uint32(42).string(message.rider);
    }
    if (message.riderWager !== "") {
      writer.uint32(50).string(message.riderWager);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewEscrow {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgNewEscrow } as MsgNewEscrow;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.strategy = reader.string();
          break;
        case 3:
          message.instigator = reader.string();
          break;
        case 4:
          message.instigatorWager = reader.string();
          break;
        case 5:
          message.rider = reader.string();
          break;
        case 6:
          message.riderWager = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgNewEscrow {
    const message = { ...baseMsgNewEscrow } as MsgNewEscrow;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.strategy !== undefined && object.strategy !== null) {
      message.strategy = String(object.strategy);
    } else {
      message.strategy = "";
    }
    if (object.instigator !== undefined && object.instigator !== null) {
      message.instigator = String(object.instigator);
    } else {
      message.instigator = "";
    }
    if (
      object.instigatorWager !== undefined &&
      object.instigatorWager !== null
    ) {
      message.instigatorWager = String(object.instigatorWager);
    } else {
      message.instigatorWager = "";
    }
    if (object.rider !== undefined && object.rider !== null) {
      message.rider = String(object.rider);
    } else {
      message.rider = "";
    }
    if (object.riderWager !== undefined && object.riderWager !== null) {
      message.riderWager = String(object.riderWager);
    } else {
      message.riderWager = "";
    }
    return message;
  },

  toJSON(message: MsgNewEscrow): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.strategy !== undefined && (obj.strategy = message.strategy);
    message.instigator !== undefined && (obj.instigator = message.instigator);
    message.instigatorWager !== undefined &&
      (obj.instigatorWager = message.instigatorWager);
    message.rider !== undefined && (obj.rider = message.rider);
    message.riderWager !== undefined && (obj.riderWager = message.riderWager);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgNewEscrow>): MsgNewEscrow {
    const message = { ...baseMsgNewEscrow } as MsgNewEscrow;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.strategy !== undefined && object.strategy !== null) {
      message.strategy = object.strategy;
    } else {
      message.strategy = "";
    }
    if (object.instigator !== undefined && object.instigator !== null) {
      message.instigator = object.instigator;
    } else {
      message.instigator = "";
    }
    if (
      object.instigatorWager !== undefined &&
      object.instigatorWager !== null
    ) {
      message.instigatorWager = object.instigatorWager;
    } else {
      message.instigatorWager = "";
    }
    if (object.rider !== undefined && object.rider !== null) {
      message.rider = object.rider;
    } else {
      message.rider = "";
    }
    if (object.riderWager !== undefined && object.riderWager !== null) {
      message.riderWager = object.riderWager;
    } else {
      message.riderWager = "";
    }
    return message;
  },
};

const baseMsgNewEscrowResponse: object = { escrowIndex: "" };

export const MsgNewEscrowResponse = {
  encode(
    message: MsgNewEscrowResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.escrowIndex !== "") {
      writer.uint32(10).string(message.escrowIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgNewEscrowResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgNewEscrowResponse } as MsgNewEscrowResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.escrowIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgNewEscrowResponse {
    const message = { ...baseMsgNewEscrowResponse } as MsgNewEscrowResponse;
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = String(object.escrowIndex);
    } else {
      message.escrowIndex = "";
    }
    return message;
  },

  toJSON(message: MsgNewEscrowResponse): unknown {
    const obj: any = {};
    message.escrowIndex !== undefined &&
      (obj.escrowIndex = message.escrowIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgNewEscrowResponse>): MsgNewEscrowResponse {
    const message = { ...baseMsgNewEscrowResponse } as MsgNewEscrowResponse;
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = object.escrowIndex;
    } else {
      message.escrowIndex = "";
    }
    return message;
  },
};

const baseMsgApprove: object = { creator: "", escrowIndex: "" };

export const MsgApprove = {
  encode(message: MsgApprove, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.escrowIndex !== "") {
      writer.uint32(18).string(message.escrowIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApprove {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApprove } as MsgApprove;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.escrowIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApprove {
    const message = { ...baseMsgApprove } as MsgApprove;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = String(object.escrowIndex);
    } else {
      message.escrowIndex = "";
    }
    return message;
  },

  toJSON(message: MsgApprove): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.escrowIndex !== undefined &&
      (obj.escrowIndex = message.escrowIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgApprove>): MsgApprove {
    const message = { ...baseMsgApprove } as MsgApprove;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = object.escrowIndex;
    } else {
      message.escrowIndex = "";
    }
    return message;
  },
};

const baseMsgApproveResponse: object = {};

export const MsgApproveResponse = {
  encode(_: MsgApproveResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveResponse } as MsgApproveResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgApproveResponse {
    const message = { ...baseMsgApproveResponse } as MsgApproveResponse;
    return message;
  },

  toJSON(_: MsgApproveResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgApproveResponse>): MsgApproveResponse {
    const message = { ...baseMsgApproveResponse } as MsgApproveResponse;
    return message;
  },
};

const baseMsgDeposit: object = { creator: "", escrowIndex: "" };

export const MsgDeposit = {
  encode(message: MsgDeposit, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.escrowIndex !== "") {
      writer.uint32(18).string(message.escrowIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeposit {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeposit } as MsgDeposit;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.escrowIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = String(object.escrowIndex);
    } else {
      message.escrowIndex = "";
    }
    return message;
  },

  toJSON(message: MsgDeposit): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.escrowIndex !== undefined &&
      (obj.escrowIndex = message.escrowIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeposit>): MsgDeposit {
    const message = { ...baseMsgDeposit } as MsgDeposit;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = object.escrowIndex;
    } else {
      message.escrowIndex = "";
    }
    return message;
  },
};

const baseMsgDepositResponse: object = {};

export const MsgDepositResponse = {
  encode(_: MsgDepositResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDepositResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    return message;
  },

  toJSON(_: MsgDepositResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgDepositResponse>): MsgDepositResponse {
    const message = { ...baseMsgDepositResponse } as MsgDepositResponse;
    return message;
  },
};

const baseMsgConfirm: object = { creator: "", escrowIndex: "", winner: "" };

export const MsgConfirm = {
  encode(message: MsgConfirm, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.escrowIndex !== "") {
      writer.uint32(18).string(message.escrowIndex);
    }
    if (message.winner !== "") {
      writer.uint32(26).string(message.winner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfirm {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConfirm } as MsgConfirm;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.escrowIndex = reader.string();
          break;
        case 3:
          message.winner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgConfirm {
    const message = { ...baseMsgConfirm } as MsgConfirm;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = String(object.escrowIndex);
    } else {
      message.escrowIndex = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
    }
    return message;
  },

  toJSON(message: MsgConfirm): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.escrowIndex !== undefined &&
      (obj.escrowIndex = message.escrowIndex);
    message.winner !== undefined && (obj.winner = message.winner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgConfirm>): MsgConfirm {
    const message = { ...baseMsgConfirm } as MsgConfirm;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = object.escrowIndex;
    } else {
      message.escrowIndex = "";
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
    }
    return message;
  },
};

const baseMsgConfirmResponse: object = {};

export const MsgConfirmResponse = {
  encode(_: MsgConfirmResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgConfirmResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgConfirmResponse } as MsgConfirmResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgConfirmResponse {
    const message = { ...baseMsgConfirmResponse } as MsgConfirmResponse;
    return message;
  },

  toJSON(_: MsgConfirmResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgConfirmResponse>): MsgConfirmResponse {
    const message = { ...baseMsgConfirmResponse } as MsgConfirmResponse;
    return message;
  },
};

const baseMsgWithdraw: object = { creator: "", escrowIndex: "" };

export const MsgWithdraw = {
  encode(message: MsgWithdraw, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.escrowIndex !== "") {
      writer.uint32(18).string(message.escrowIndex);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdraw {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.escrowIndex = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgWithdraw {
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = String(object.escrowIndex);
    } else {
      message.escrowIndex = "";
    }
    return message;
  },

  toJSON(message: MsgWithdraw): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.escrowIndex !== undefined &&
      (obj.escrowIndex = message.escrowIndex);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgWithdraw>): MsgWithdraw {
    const message = { ...baseMsgWithdraw } as MsgWithdraw;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.escrowIndex !== undefined && object.escrowIndex !== null) {
      message.escrowIndex = object.escrowIndex;
    } else {
      message.escrowIndex = "";
    }
    return message;
  },
};

const baseMsgWithdrawResponse: object = {};

export const MsgWithdrawResponse = {
  encode(_: MsgWithdrawResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgWithdrawResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgWithdrawResponse {
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    return message;
  },

  toJSON(_: MsgWithdrawResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgWithdrawResponse>): MsgWithdrawResponse {
    const message = { ...baseMsgWithdrawResponse } as MsgWithdrawResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  NewEscrow(request: MsgNewEscrow): Promise<MsgNewEscrowResponse>;
  Approve(request: MsgApprove): Promise<MsgApproveResponse>;
  Deposit(request: MsgDeposit): Promise<MsgDepositResponse>;
  Confirm(request: MsgConfirm): Promise<MsgConfirmResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Withdraw(request: MsgWithdraw): Promise<MsgWithdrawResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  NewEscrow(request: MsgNewEscrow): Promise<MsgNewEscrowResponse> {
    const data = MsgNewEscrow.encode(request).finish();
    const promise = this.rpc.request("escrow.escrow.Msg", "NewEscrow", data);
    return promise.then((data) =>
      MsgNewEscrowResponse.decode(new Reader(data))
    );
  }

  Approve(request: MsgApprove): Promise<MsgApproveResponse> {
    const data = MsgApprove.encode(request).finish();
    const promise = this.rpc.request("escrow.escrow.Msg", "Approve", data);
    return promise.then((data) => MsgApproveResponse.decode(new Reader(data)));
  }

  Deposit(request: MsgDeposit): Promise<MsgDepositResponse> {
    const data = MsgDeposit.encode(request).finish();
    const promise = this.rpc.request("escrow.escrow.Msg", "Deposit", data);
    return promise.then((data) => MsgDepositResponse.decode(new Reader(data)));
  }

  Confirm(request: MsgConfirm): Promise<MsgConfirmResponse> {
    const data = MsgConfirm.encode(request).finish();
    const promise = this.rpc.request("escrow.escrow.Msg", "Confirm", data);
    return promise.then((data) => MsgConfirmResponse.decode(new Reader(data)));
  }

  Withdraw(request: MsgWithdraw): Promise<MsgWithdrawResponse> {
    const data = MsgWithdraw.encode(request).finish();
    const promise = this.rpc.request("escrow.escrow.Msg", "Withdraw", data);
    return promise.then((data) => MsgWithdrawResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
