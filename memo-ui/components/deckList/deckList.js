import { Flex, Box } from "rebass"
import Deck from "../deck/deck"

function DeckList({ decks }) {
    console.log(decks)
    return (
        <Flex>

            {decks.map((deck, index) => {
                return (
                    <Box justifyContent='center'
                        p={3}
                        width={1 / 2}
                        bg='primary'>
                        <Deck {...deck} />
                    </Box>
                )
            })}
        </Flex>
    )
}

export default DeckList

