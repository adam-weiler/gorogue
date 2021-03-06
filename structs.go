/*
Copyright (c) 2018 Tomasz "VedVid" Nowakowski

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package main

type BasicProperties struct {
	/*BasicProperties is struct that aggregates
	all widely used data, necessary for every
	map tile and object representation*/
	Layer     int
	X, Y      int
	Char      string
	Color     string
	ColorDark string
}

type VisibilityProperties struct {
	/*VisibilityProperties is simple struct
	for checking if object is always visible,
	regardless of player's fov*/
	AlwaysVisible bool
}

type CollisionProperties struct {
	/*CollisionProperties is struct filled with
	boolean values, for checking several
	collision conditions: if cell is blocked,
	if it blocks creature sight, etc.*/
	Blocked     bool
	BlocksSight bool
}

type AIProperties struct {
	/*AIProperties serves holding information about
	monster AI. AI types are iota (integers) in
	monsters.go*/
	AIType int
}
