import * as React from "react";

export default function Labels({
  connectedText,
  rollsText,
  uptimeText,
}: {
  connectedText: string;
  rollsText: string;
  uptimeText: string;
}) {
  return (
    <m-group y="7">
      <m-label
        id="uptime-label"
        y="2"
        width="4"
        height="0.5"
        color="#bfdbfe"
        font-color="#172554"
        alignment="center"
        content={uptimeText}
      ></m-label>
      <m-label
        id="connected-label"
        content={connectedText}
        y="1"
        width="4"
        height="0.5"
        color="#bfdbfe"
        font-color="#172554"
        alignment="center"
      ></m-label>
      <m-label
        id="click-label"
        y="0"
        width="4"
        height="0.5"
        color="#bfdbfe"
        font-color="#172554"
        alignment="center"
        content={rollsText}
      ></m-label>
    </m-group>
  );
}
