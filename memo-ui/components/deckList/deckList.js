import { Flex } from "rebass"
import Deck from "../deck/deck"

function DeckList({decks}) {
    console.log(decks)
    return (
        <Flex>
            {decks.map((deck, index) => {
              return <Deck {...deck}/>
            })}
        </Flex>
    )
}

export default DeckList