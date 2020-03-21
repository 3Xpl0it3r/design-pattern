package main

import "design-pattern/behavioral/command"

func main() {
	//standardMaze := new(creational.MazeGame).CreateMazeFromAbstractFactory(new(creational.MazeStandardAbstractFactory))
	//standardMaze.Enter()

	//enchantedMaze := new(creational.MazeGame).CreateMazeFromAbstractFactory(new(creational.MazeEnchantedAbstractFactory))
	//enchantedMaze.Enter()

	//buildrrMaze := new(creational.MazeGame).CreateMazeFromBuilder(new(creational.MazeStandardBuilder))
	//buildrrMaze.Enter()
	//
	//factroyMethodMze := new(creational.StandardMazeFactoryMethod).CreateMaze()
	//factroyMethodMze.Enter()
	//

	//var mazeGame *creational.MazeGame
	//
	//singletonFactory := creational.GetMazeFactory()
	//maze := mazeGame.CreateMazeFromAbstractFactory(singletonFactory.Factory)
	//maze.Enter()

	//
	//standardFactory := new(creational.MazeStandardAbstractFactory)
	//enchantedFactory := new(creational.MazeEnchantedAbstractFactory)
	//
	//creational.RegisterFactoryToSingleton("standard", standardFactory)
	//creational.RegisterFactoryToSingleton("enchanted", enchantedFactory)
	//
	//standardMaze := new(creational.MazeGame).CreateMazeFromAbstractFactory(creational.GetFactoryByName("standard"))
	//standardMaze.Enter()
	//
	//enchantedMaze := new(creational.MazeGame).CreateMazeFromAbstractFactory(creational.GetFactoryByName("enchanted"))
	//enchantedMaze.Enter()

	invoker := command.NewInvoker()
	cmdA := command.NewConcreteCommandA()
	cmdB := command.NewConcreteCommandB()

	recA := command.NewReceiver()
	recB := command.NewReceiverB()

	cmdB.SetReceiver(*recB)
	cmdA.SetReceiver(*recA)

	invoker.AddCommand(cmdA)
	invoker.AddCommand(cmdB)
	invoker.ExecuteCommand()


}
