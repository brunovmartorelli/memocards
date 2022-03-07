import { Flex, Box, Link, Button } from "rebass"
import theme from "../../theme"
import Deck from "../deck/deck"

function DeckList({ decks }) {
    console.log(decks)
    return (
        <Flex marginY='30px'>

            {decks.map((deck, index) => {
                return (
                    <Flex 
                        justifyContent="center"
                        flexDirection="column"
                        color="#FFF"
                        style={{
                            border: '1px solid #FFF',
                            margin: '0 10px'
                        }}
                        p={10}
                    >
                        <Deck {...deck} />
                        <Flex>
                            <Link href={`/decks/${deck.name}`}
                                color={theme.lightgrey}
                                justifyContent='center'
                                p={3}
                                width={1 / 2}
                                bg='primary'
                            >
                                <Button fontSize={2}
                                    style={{ cursor: 'pointer' }}
                                    backgroundColor={theme.darkgrey}
                                    fontFamily='Roboto'>
                                    Editar
                                </Button>
                            </Link>

                            <Link href={`/decks/${deck.name}/study`}
                                color={theme.lightgrey}
                                justifyContent='center'
                                p={3}
                                width={1 / 2}
                                bg='primary'
                            >
                                <Button fontSize={2}
                                    style={{ cursor: 'pointer' }}
                                    backgroundColor={theme.discordblue}
                                    fontFamily='Roboto'>
                                    Estudar
                                </Button>
                            </Link>
                        </Flex>
                    </Flex>
                )
            })}
        </Flex>
    )
}

export default DeckList

