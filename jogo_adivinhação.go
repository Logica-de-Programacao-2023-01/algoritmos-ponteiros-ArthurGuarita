package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print("Bem-vindo ao jogo da adivinhação!\n\n")
	//O programa deve gerar um número aleatório entre 1 e 100 que será a resposta
	var resposta = rand.Intn(100) + 1 // add 1 devido ao intervalo 0 a 99
	//
	var adivinhacao int
	//
	var tentativasPorJogada []int
	//
	var jogarNovamente string
	//
	var round, soma, count = 0, 0, 0
	//
	for {
		fmt.Print("Digite um valor inteiro entre 1 e 100: ")
		_, err := fmt.Scan(&adivinhacao)
		if err != nil {
			fmt.Print("Ocorreu um erro")
			continue
		}
		if adivinhacao != resposta {
			comparar(adivinhacao, resposta)
			count++
			continue
		} else {
			count++
			fmt.Print("Parabéns, você acertou!\n")
			fmt.Printf("Você utilizou %d tentativas\n", count)
			tentativasPorJogada = append(tentativasPorJogada, count)

			//jogar de novo

			fmt.Print("Deseja jogar novamente(s/n)?: ")
			_, er := fmt.Scan(&jogarNovamente)
			if er != nil {
				fmt.Printf("ocorreu um erro %v", er)
				return
			}
			if jogarNovamente == "s" {
				resposta = rand.Intn(100) + 1
				count = 0
				continue
			} else {
				for _, t := range tentativasPorJogada {
					round++
					soma += t
					fmt.Printf("Você utilizou %v tentativas no round %v\n", t, round)
				}
				fmt.Print("O total de tentativas em todas as jogadas foi: ", soma)
				return
			}

		}
	}

}

func comparar(adivinhacao, resposta int) {
	if adivinhacao > resposta {
		fmt.Printf("O número é menor que %d\n", adivinhacao)
	} else if adivinhacao < resposta {
		fmt.Printf("O número é maior que %d\n", adivinhacao)
	}
}
