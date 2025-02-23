import type { Address } from "viem"

export type Source = {
  sourceAddress: Address
  owner: Address
}

export type CustomContract = {
  contractAddress: Address
  addedBy: Address
  contractABI: string
}

export type UseEtherbaseSourceProps = Readonly<{
  sourceAddress: Address
}>

export type UseEtherbaseContractProps = Readonly<{
  contractAddress: Address
}>

export type UseEtherbasePermissionsProps = Readonly<{
  sourceAddress: Address
}>
