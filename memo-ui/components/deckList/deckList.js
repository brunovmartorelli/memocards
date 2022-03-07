import { Flex, Box, Link, Button } from "rebass"
import { listStudy } from "../../services/card.service"
import theme from "../../theme"
import Deck from "../deck/deck"
import { useRouter } from "next/router"

function DeckList({ decks }) {
    const router = useRouter();

    const onEstudarClick = async (name) => {
        const cards = await listStudy(name)
        
        if (cards) {
            router.push(`/decks/${name}/study`)
        } else {
            alert("Não há mais cartas para estudar no momento.")
        }
    }
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

                            <Button
                                onClick={() => onEstudarClick(deck.name)}
                                color={theme.lightgrey}
                                justifyContent='center'
                                bg='primary'
                                height={35}
                                mt={3}
                                fontSize={2}
                                style={{ cursor: 'pointer' }}
                                backgroundColor={theme.discordblue}
                                fontFamily='Roboto'>
                                Estudar
                            </Button>
                        </Flex>
                    </Flex>
                )
            })}
        </Flex>
    )
}

export default DeckList

