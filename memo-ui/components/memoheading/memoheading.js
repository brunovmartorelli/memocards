import { Flex, Box, Heading, Button } from "rebass"
import Link from "next/link"
import theme from "../../theme"
import { useRouter } from "next/router"

function MemoHeading({ title, color }) {
    const router = useRouter();
    const notHome = router.asPath !== "/"
    const defaultColor = 'white'
    return (
        <Heading
            fontFamily='Roboto'
            fontSize={[7]}
            color={color ? color : defaultColor}
        >
            { notHome && <Flex>
                <Box px='auto'>
                    <Link href='/'>
                        <Button fontSize={5}
                            style={{ cursor: 'pointer' }}
                            backgroundColor={theme.darkgrey}
                            fontFamily='Roboto'>
                            Home
                        </Button>
                    </Link>
                </Box>
            </Flex>}
            <Flex justifyContent='center'>
                <Box px='auto'>
                    {title}
                </Box>
            </Flex>
            <hr />
        </Heading>
    )
}

export default MemoHeading