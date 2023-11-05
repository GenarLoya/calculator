import React from "react";
import CalculatorButton from "./buttonRender";
import { Toaster, toast } from "sonner";

export default function CalcRender() {
  const [display, setDisplay] = React.useState("");

  function deleteLastChar() {
    console.log("deleteLastChar");
    setDisplay(display.slice(0, -1));
  }

  function clear() {
    console.log("clear");
    setDisplay("");
  }

  function placeSymbol(symbol = "") {
    if (symbol === "") return;
    console.log("placeSymbol", symbol);
    setDisplay(display + symbol);
  }

  function calculate() {
    console.log("calculate");
    console.log(display);

    fetch("http://localhost:8080/calculate", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ exp: display }),
    })
      .then((res) => res.json())
      .then((data) => {
        setDisplay(data.result);
        toast.success("Calculation successful");
      })
      .catch((err) => {
        toast.error("Calculation failed:" + err);
      });
  }

  return (
    <div
      role="calc-body"
      className="flex flex-col p-2 w-max h-max rounded-3 bg-slate-6 shadow-lg shadow-inner"
    >
      <Toaster />
      <span className="h-max flex justify-end text-sm">CASIO</span>
      <div className="flex space-x-2 h-10%">
        <input
          type="text"
          value={display}
          role="calc-display"
          className="h-full px-2 pt-3 mb-5  bg-neutral-3 rounded-3 text-lg cursor-cell ring-0 focus:ring-0 focus:shadow-slate-8 focus:shadow-lg focus:outline-none duration-150 border-0"
          style={{ fontFamily: "system-ui" }}
          onChange={(e) => setDisplay(e.target.value)}
        ></input>

        <button className="bg-red rounded-3 h-fullpx-2 pt-3 mb-5 border-none focus:outline-blueGray duration-150 hover:scale-150" onClick={deleteLastChar}>Del</button>
      </div>
      <section role="calc-buttons" className="grid grid-cols-4 gap-4 px-1 py-2">
        <CalculatorButton value="C" onClick={clear} />
        <CalculatorButton value="^" onClick={placeSymbol} columns={2} />
        <CalculatorButton value="/" onClick={placeSymbol} />
        <CalculatorButton value="7" onClick={placeSymbol} />
        <CalculatorButton value="8" onClick={placeSymbol} />
        <CalculatorButton value="9" onClick={placeSymbol} />
        <CalculatorButton value="*" onClick={placeSymbol} />
        <CalculatorButton value="4" onClick={placeSymbol} />
        <CalculatorButton value="5" onClick={placeSymbol} />
        <CalculatorButton value="6" onClick={placeSymbol} />
        <CalculatorButton value="-" onClick={placeSymbol} />
        <CalculatorButton value="1" onClick={placeSymbol} />
        <CalculatorButton value="2" onClick={placeSymbol} />
        <CalculatorButton value="3" onClick={placeSymbol} />
        <CalculatorButton value="+" onClick={placeSymbol} />
        <CalculatorButton value="(" onClick={placeSymbol} />
        <CalculatorButton value="0" onClick={placeSymbol} />
        <CalculatorButton value=")" onClick={placeSymbol} />
        <CalculatorButton value="=" onClick={calculate} />
      </section>
    </div>
  );
}
