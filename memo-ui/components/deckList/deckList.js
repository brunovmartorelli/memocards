import { Flex, Box } from "rebass"
import theme from "../../theme"
import Deck from "../deck/deck"

function DeckList({ decks }) {
    console.log(decks)
    return (
        <Flex>

            {decks.map((deck, index) => {
                return (
                    <Box
                        color={theme.lightgrey}
                        justifyContent='center'
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

