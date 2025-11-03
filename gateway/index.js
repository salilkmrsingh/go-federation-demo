import { ApolloServer } from '@apollo/server';
import { ApolloGateway } from '@apollo/gateway';
import { startStandaloneServer } from '@apollo/server/standalone';

const gateway = new ApolloGateway({
    serviceList: [
        { name: 'users', url: 'http://localhost:8080/query' },
        { name: 'products', url: 'http://localhost:8081/query' },
        { name: 'reviews', url: 'http://localhost:8082/query' },
    ],
});

const server = new ApolloServer({
    gateway,
    includeStacktraceInErrorResponses: false,

    formatError: (formattedError) => {
        // Remove serviceName from the extensions
        if (formattedError.extensions) {
            const { serviceName, ...rest } = formattedError.extensions;
            formattedError.extensions = rest;
        }
        return formattedError;
    },
});
const { url } = await startStandaloneServer(server, { listen: { port: 8085 } });

console.log(`🚀 Gateway ready at ${url}`);
