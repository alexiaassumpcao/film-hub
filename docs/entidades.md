# Hub de filmes

## Planejamento de entidades:

### filme
_id: uui
nome: string
categoria: [string]
ranking: float
data de estreia: date
tempo de duração: number(em minutos)
constinuações: [string]
=> Exemplo:
{
	_id: 56sdcjsnc1212,
	name: "A Escolha Perfeita",
	category: ["musical", "comedy"] ,
	ranking: 4.5,
	release_date: 09-24-2012,
	running_time: 122,
	continuation: ["A Escolha perfeita 2", "A Escolha perfeita 3"]
}

### usuário: 
_id: uui
nome: string
questão de acesso: string
resposta: string
filmes favoritos: [string]
categoria favorita: [string]
=> Exemplo:
{
	_id: 5864snjnieunve5668,
	name: Jobson Pedreiro,
	access_question: "How many friends do i have?",
	access_answer: "5",
	favorites: {
		films: ["A escolha Perfeita"],
		category: ["musical", "terror"]
	}
}

### categoria
_id: uui
nome: string
caracteristicas: string
films:[string]
=> Exemplo:
{
	_id: 593ngnviojrovm518,
	name: "comedy",
	features: "Filmes que provocam risadas",
	films: ["A escolha perfeita"]
}