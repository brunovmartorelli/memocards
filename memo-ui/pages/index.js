import DeckList from "../components/deckList/deckList"
import { list } from "./../services/deck.service"
import { ThemeProvider, withTheme } from '@emotion/react'
import theme from '@rebass/preset'
import { Box, Flex, Heading } from "rebass"

const Index = ({ decks }) => {
    return (
        <ThemeProvider theme={theme}>
            <Heading
                fontSize={[7]}
                color='blue'
            >
                <Flex justifyContent='center'>
                    <Box px='auto'>
                        MEMOCARDS
                    </Box>
                </Flex>
            </Heading>
            <DeckList decks={decks} />
        </ThemeProvider>
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