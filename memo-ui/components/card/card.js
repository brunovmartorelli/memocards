import { Flex, Text } from "rebass"

const Card = ({ front, back, score, ReviewedAt }) => {
    return (
        <Flex flexDirection="column">
            <Text fontSize={[5]}
                fontWeight='bold'
                textAlign='center'>Front: {front}</Text>
            <Text fontSize={[5]}
                fontWeight='bold'
                textAlign='center'>Back: {back}</Text>
            <Text fontSize={[5]}
                fontWeight='bold'
                textAlign='center'>Score: {score}</Text>

        </Flex>
    )
}

export default Card;