import DeckList from "../components/deckList/deckList"
import { list } from "./../services/deck.service"
import { withTheme } from '@emotion/react'
import MemoHeading from "./../components/memoheading/memoheading"
import { Flex, Box, Button } from "rebass"
import theme from "../theme"
import Link from "next/link"

const Index = ({ decks }) => {
    return (
        <>
            <MemoHeading title='MEMOCARDS' />
            <DeckList decks={decks} />
            <Flex justifyContent='center' mx={-2} mb={3}>
                <Box px={2}>
                    <Link href='/createDeck'>
                        <Button
                            style={{ cursor: 'pointer' }}
                            backgroundColor={theme.discordblue}
                            fontFamily='Roboto'>
                            Criar Deck
                        </Button>
                    </Link>
                </Box>
            </Flex>
        </>
    )
}

export async function getStaticProps() {
    const decks = await list()
    return {
        props: {
            decks: decks
        }
    }
}



export default withTheme(Index)