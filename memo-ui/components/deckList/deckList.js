import { Flex, Box, Link } from "rebass"
import theme from "../../theme"
import Deck from "../deck/deck"

function DeckList({ decks }) {
    console.log(decks)
    return (
        <Flex marginY='30px'>

            {decks.map((deck, index) => {
                return (
                    <Link href={`/decks/${deck.name}`}
                        style={{ 
                            cursor: 'pointer',
                            border: '1px solid #FFF',
                            margin: '0 10px'
                        }}
                        color={theme.lightgrey}
                        justifyContent='center'
                        p={3}
                        width={1 / 2}
                        bg='primary'>
                        <Deck {...deck} />
                    </Link>
                )
            })}
        </Flex>
    )
}

export default DeckList

