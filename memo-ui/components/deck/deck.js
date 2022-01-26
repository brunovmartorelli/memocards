import { Text } from "rebass";

const Deck = ({ name, description }) => {
    return (
        <div>
            <Text
                fontSize={[5]}
                fontWeight='bold'
            >
                {name}
            </Text>
            <Text
                fontSize={[3]}
                fontWeight='bold'
                color='primary'
            >
                {description}
            </Text>
        </div>
    )
}

export default Deck;