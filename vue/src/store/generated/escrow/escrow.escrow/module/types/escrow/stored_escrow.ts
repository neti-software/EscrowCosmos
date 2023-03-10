/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "escrow.escrow";

export interface StoredEscrow {
  index: string;
  escrow: string;
}

const baseStoredEscrow: object = { index: "", escrow: "" };

export const StoredEscrow = {
  encode(message: StoredEscrow, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.escrow !== "") {
      writer.uint32(18).string(message.escrow);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredEscrow {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredEscrow } as StoredEscrow;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.escrow = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredEscrow {
    const message = { ...baseStoredEscrow } as StoredEscrow;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.escrow !== undefined && object.escrow !== null) {
      message.escrow = String(object.escrow);
    } else {
      message.escrow = "";
    }
    return message;
  },

  toJSON(message: StoredEscrow): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.escrow !== undefined && (obj.escrow = message.escrow);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredEscrow>): StoredEscrow {
    const message = { ...baseStoredEscrow } as StoredEscrow;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.escrow !== undefined && object.escrow !== null) {
      message.escrow = object.escrow;
    } else {
      message.escrow = "";
    }
    return message;
  },
};

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
