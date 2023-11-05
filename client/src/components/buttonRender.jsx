import PropTypes from "prop-types";
import classNames from "classnames";

const CalculatorButton = ({ value, onClick, columns = 1 }) => {
  return (
    <button
      onClick={() => {
        typeof onClick === "function" && onClick(value);
      }}
      className={classNames(
        "rounded-3 border-none border-gray-500 bg-slate-300 p-4 align-middle duration-150 hover:scale-130 hover:shadow-lg hover:shadow-slate-500 hover:bg-slate-400 hover:text-white hover:font-bold hover:rounded-3 hover:border-none hover:border-gray-500 hover:bg-slate-300 hover:p-4 hover:align-middle",
        {
          "col-span-2": columns === 2,
          "col-span-3": columns === 3,
          "col-span-4": columns === 4,
        }
      )}
    >
      {value}
    </button>
  );
};

CalculatorButton.propTypes = {
  value: PropTypes.string.isRequired,
  onClick: PropTypes.func,
  //Check columns in max 4 nd min 1
  columns: (props, propName, componentName) =>
    [1, 2, 3, 4].includes(props[propName]) || !props[propName]
      ? null
      : new Error(`Invalid prop ${propName} on component ${componentName}`),
};

export default CalculatorButton;
