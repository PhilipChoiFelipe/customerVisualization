import React, { useState } from "react";
import { useSelector } from 'react-redux';
import {
    VictoryPie,
    VictoryBar,
    VictoryTheme,
    VictoryChart,
    VictoryVoronoiContainer,
    VictoryTooltip,
} from 'victory';
import _ from "lodash";
import { Button } from 'react-bootstrap';


/** 
* @description
[data field]ToCus convert customer data to max count
{
    [data field]: ''
    count: ''
}
*/
const ethToCus = (customers) => {
    let temp = _.countBy(customers, 'ethnicity')
    let dataToViz = []
    for (let key in temp) {
        dataToViz.push({
            ethnicity: key,
            count: temp[key]
        })
    }
    return (
        <VictoryPie
            data={dataToViz}
            labels={({ datum }) => `${datum.ethnicity}: ${datum.count}`}
            width={600}
            x="ethnicity"
            y="count"
            colorScale={["tomato", "orange", "gold", "cyan", "navy"]}
            cornerRadius={({ datum }) => datum.count}
        />
    );
}

const disToCus = (customers) => {
    let temp = _.countBy(customers, 'disChannel')
    let dataToViz = []
    for (let key in temp) {
        dataToViz.push({
            disChannel: key,
            count: temp[key]
        })
    }
    return (
        <VictoryChart
            theme={VictoryTheme.material}
            containerComponent={<VictoryVoronoiContainer />}
            width={600}
        >
            <VictoryBar
                data={dataToViz}
                label="disChannel"
                x="disChannel"
                y="count"
                labelRadius={({ innerRadius }) => innerRadius + 50}
                labelComponent={<VictoryTooltip />}
                containerComponent={<VictoryVoronoiContainer />}
            />
        </VictoryChart>
    );
}

const gendToCus = (customers) => {

    let temp = _.countBy(customers, 'gender')
    let dataToViz = []
    for (let key in temp) {
        dataToViz.push({
            gender: key,
            count: temp[key]
        })
    }
    return (
        <VictoryPie
            data={dataToViz}
            width={600}
            labels={({ datum }) => `${datum.gender}: ${datum.count}`}
            x="gender"
            y="count"
            colorScale={["tomato", "orange", "gold", "cyan", "navy"]}
            cornerRadius={({ datum }) => datum.count}
        />
    );
}

/**
 *@description Visualization visualize customers with different data field
*/
const Visualization = () => {
    const [visType, setVisType] = useState(null);

    const { customers } = useSelector((state) => state.customer);
    return (
        <div className="container">
            <h1>
                Visualize your customers type
        </h1>
            <Button variant="success" onClick={() => setVisType(() => gendToCus(customers))}>Gender</Button>
            <Button variant="warning" onClick={() => setVisType(() => disToCus(customers))}>Discovered Channel</Button>
            <Button variant="info" onClick={() => setVisType(() => ethToCus(customers))}>Ethnicity</Button>
            {visType}
        </div>
    );
}

export default Visualization;