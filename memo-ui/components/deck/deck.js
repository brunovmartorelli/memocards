import { Text, Box } from "rebass";



const Deck = ({ name, description }) => {
    return (
        <Box justifyContent='center'>
            <Text
                fontSize={[5]}
                fontWeight='bold'
                textAlign='center'
            >
                {name}
            </Text>
            <Text
                fontSize={[3]}
                fontWeight='bold'
                textAlign='center'
            >
                {description}
            </Text>
        </Box>
    )
}

export default Deck;