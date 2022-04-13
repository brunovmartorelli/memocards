import { Flex, Box, Link, Button } from "rebass"
import theme from "../../theme"
import Card from "../card/card"
import { deleteCard } from '../../services/card.service'
import { useRouter } from "next/router"

function CardList({ cards, deckName }) {
    const router = useRouter();

    const onDeletarClick = async (front) => {
        await deleteCard(front, deckName)
        router.push(`/decks/${deckName}`)
    }
    return (
        <Flex marginY='30px'>

            {cards.map((card, index) => {
                return (
                    <Box
                        style={{
                            border: '1px solid #FFF',
                            margin: '0 10px'
                        }}
                        color={theme.lightgrey}
                        justifyContent='center'
                        textAlign='center'
                        p={3}
                        width={1 / 2}
                        bg='primary'>
                        <Card {...card} />
                        <Button
                            alignSelf='center'
                            onClick={() => onDeletarClick(card.front)}
                            height={35}
                            mt={3}
                            ml={4}
                            fontSize={2}
                            style={{ cursor: 'pointer' }}
                            backgroundColor={theme.red}
                            fontFamily='Roboto'>
                            Deletar
                        </Button>
                    </Box>
                )
            })}
        </Flex>
    )
}

export default CardList