let isEnvProduction = (process.argv[2] === "production")
export const SCHEME   = (isEnvProduction ? "https://"    : "http://")
export const DOMAIN   = "localhost:3000"
export const ENDPOINT = ""

export const PATH = {
  CREATE_TICKET: '/backlog/ticket'
}
