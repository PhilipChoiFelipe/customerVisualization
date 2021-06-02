import React from "react";
import { useSelector } from 'react-redux';
import {
    VictoryChart,
    VictoryAxis,
    VictoryBar,
    VictoryStack
} from 'victory';

const Visualization = () => {
    const { customers } = useSelector((state) => state.customer);
    console.log(customers);
    return(
        <>
        <h1>
            Visualize your customers type
        </h1>
        <VictoryBar/>
        </>
    );
}

export default Visualization;