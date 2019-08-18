import { test, runTests } from "https://deno.land/std/testing/mod.ts";
import { assertEquals } from "https://deno.land/std/testing/asserts.ts";

import { start, State, Mode, Point } from './main.ts';

test(function t1() {
  assertEquals("hello", "hello");
});

test(function t2() {
  assertEquals("world", "world");
});

test(function testStart() {
  assertEquals(start, {Mode: Mode.Up, Point: {x: 0, y: 0}});
});

runTests();
