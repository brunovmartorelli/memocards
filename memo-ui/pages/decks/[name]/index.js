import MemoHeading from "../../../components/memoheading/memoheading"
import { list } from '../../../services/card.service'
import CardList from "../../../components/cardList/cardList"
import { Flex, Box, Button } from "rebass"
import Link from "next/link"
import theme from "../../../theme"

function DeckPage({cards, name}) {

    return (
        <>
            <MemoHeading title={name} />
            <Flex justifyContent='center' mx={-2} mb={3}>
                <Box px={2}>
                    <Link href={`/decks/${name}/createCard`}>
                        <Button
                            style={{ cursor: 'pointer' }}
                            backgroundColor={theme.discordblue}
                            fontFamily='Roboto'>
                            Criar Carta
                        </Button>
                    </Link>
                </Box>
            </Flex>
            <CardList cards={cards} />
        </>
    )
}

export async function getServerSideProps({query}) {
    const name = query.name
    const cards = await list(name)
    return {
        props: {
            name: name,
            cards: cards
        }
    }
}

export default DeckPage