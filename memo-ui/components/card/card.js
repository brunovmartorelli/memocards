import { Flex, Text } from "rebass"

const Card = ({ front, back, score, reviewedAt }) => {
    const date = new Date(reviewedAt);
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
             <Text fontSize={[5]}
                fontWeight='bold'
                textAlign='center'>Data:{date.getDate()}/{date.getMonth() + 1}/{date.getFullYear()}</Text>
        </Flex>
    )
}

export default Card;