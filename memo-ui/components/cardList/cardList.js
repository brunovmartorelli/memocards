import { Flex, Box, Link } from "rebass"
import theme from "../../theme"
import Card from "../card/card"


function CardList({ cards }) {
    console.log(cards)
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
                    p={3}
                    width={1 / 2}
                    bg='primary'>
                    <Card {...card} />
                </Box>
            )
        })}
    </Flex>
    )
}

export default CardList