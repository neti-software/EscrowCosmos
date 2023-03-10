// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgNewEscrow } from "./types/escrow/tx";
import { MsgConfirm } from "./types/escrow/tx";
import { MsgApprove } from "./types/escrow/tx";
import { MsgDeposit } from "./types/escrow/tx";
import { MsgWithdraw } from "./types/escrow/tx";


const types = [
  ["/escrow.escrow.MsgNewEscrow", MsgNewEscrow],
  ["/escrow.escrow.MsgConfirm", MsgConfirm],
  ["/escrow.escrow.MsgApprove", MsgApprove],
  ["/escrow.escrow.MsgDeposit", MsgDeposit],
  ["/escrow.escrow.MsgWithdraw", MsgWithdraw],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgNewEscrow: (data: MsgNewEscrow): EncodeObject => ({ typeUrl: "/escrow.escrow.MsgNewEscrow", value: MsgNewEscrow.fromPartial( data ) }),
    msgConfirm: (data: MsgConfirm): EncodeObject => ({ typeUrl: "/escrow.escrow.MsgConfirm", value: MsgConfirm.fromPartial( data ) }),
    msgApprove: (data: MsgApprove): EncodeObject => ({ typeUrl: "/escrow.escrow.MsgApprove", value: MsgApprove.fromPartial( data ) }),
    msgDeposit: (data: MsgDeposit): EncodeObject => ({ typeUrl: "/escrow.escrow.MsgDeposit", value: MsgDeposit.fromPartial( data ) }),
    msgWithdraw: (data: MsgWithdraw): EncodeObject => ({ typeUrl: "/escrow.escrow.MsgWithdraw", value: MsgWithdraw.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
