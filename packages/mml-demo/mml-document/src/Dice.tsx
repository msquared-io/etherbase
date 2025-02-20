import { MGroupAttributes } from "@mml-io/mml-react-types"
import React, { useState } from "react"
import { useEtherbaseSource } from "@msquared/etherbase-client"

const rollMap = {
  1: [0, 0, 0],
  2: [90, 0, 0],
  3: [0, 0, 270],
  4: [0, 0, 90],
  5: [270, 0, 0],
  6: [180, 0, 0],
}
type DiceNumber = keyof typeof rollMap
const rollDuration = 400
const upDuration = 200
const downDuration = 500
const heightGain = 3
const diceHeight = 2
const halfDiceHeight = diceHeight / 2
const totalDuration = Math.max(upDuration + downDuration, rollDuration)

export default function Dice(props: MGroupAttributes) {
  const [previousResult, setPreviousResult] = useState<DiceNumber>(1)
  const [currentResult, setCurrentResult] = useState<DiceNumber>(1)
  const [rollTime, setRollTime] = useState(-totalDuration)

  const sourceAddress = "0x1be5f11a7a7e6C78ddE5aF8DaDf684049d1d5907"
  const { emitEvent } = useEtherbaseSource({ sourceAddress })

  const rollDice = async () => {
    const t = document.timeline.currentTime as number
    if (t < rollTime + totalDuration) {
      return
    }
    setRollTime(t)
    setPreviousResult(currentResult)
    const newResult = Math.floor(Math.random() * 6) + 1
    setCurrentResult(newResult as DiceNumber)
    await emitEvent({
      name: "DiceRoll",
      args: {
        result: newResult,
      },
    })
  }

  const oldRotation = rollMap[previousResult]
  const targetRotation = rollMap[currentResult]
  return (
    <m-group {...props}>
      <m-model
        id="dice"
        src="https://public.mml.io/dice.glb"
        onClick={rollDice}
      >
        <m-attr-anim
          easing="easeOutSine"
          attr="y"
          duration={upDuration}
          start-time={rollTime}
          start={halfDiceHeight.toString(10)}
          end={(heightGain + halfDiceHeight).toString(10)}
          loop={"false"}
        ></m-attr-anim>
        <m-attr-anim
          easing="easeOutBounce"
          attr="y"
          start-time={rollTime + upDuration}
          duration={downDuration}
          start={(heightGain + halfDiceHeight).toString(10)}
          end={halfDiceHeight.toString(10)}
          loop={"false"}
        ></m-attr-anim>
        <m-attr-anim
          easing="easeInOutCubic"
          attr="rx"
          duration={rollDuration}
          start-time={rollTime}
          start={oldRotation[0].toString(10)}
          end={targetRotation[0].toString(10)}
          loop={"false"}
        ></m-attr-anim>
        <m-attr-anim
          easing="easeInOutCubic"
          attr="ry"
          duration={rollDuration}
          start-time={rollTime}
          start={oldRotation[1].toString(10)}
          end={targetRotation[1].toString(10)}
          loop={"false"}
        ></m-attr-anim>
        <m-attr-anim
          easing="easeInOutCubic"
          attr="rz"
          duration={rollDuration}
          start-time={rollTime}
          start={oldRotation[2].toString(10)}
          end={targetRotation[2].toString(10)}
          loop={"false"}
        ></m-attr-anim>
      </m-model>
    </m-group>
  )
}
