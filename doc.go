// Package semdiffstat calculates semantic diffstats.
// Semantic diffstats are intended to present readable, meaningful
// high level summaries of changes to a human.
//
// TODOs:
//   * rename detection
//   * accept directories (or more), instead of going file-by-file?
//   * order insensitive diff (simply moving a function around should not generate a diff at all)
//   * handle other decl types (vars, types, consts)
//   * associate comments with their decls, at least when obvious/easy
//   * detect comment changes vs whitespace changes vs functional changes
//   * support other languages
//   * maybe provide sorting helpers: sort by name, sort by magnitude of change, sort ins-then-del-then-mod
//   * add a screenshot to the readme :P
package semdiffstat
