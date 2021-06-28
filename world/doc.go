// Package world provides the two key components for building
// simulations: worlds and agents.
//
// A world is a graph with each node assigned a context. An agent lives
// in a world and can explore it by wandering on random walks, or by moving
// through the world between his points of interest (addresses).
//
// When moving between his addresses each agent may take a slightly different route
// chosen from k shortest routes.
package world
