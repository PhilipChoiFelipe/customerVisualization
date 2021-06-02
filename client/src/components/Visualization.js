import React, {useState} from "react";
import { useSelector } from 'react-redux';
import Customers from "../data/customers.json";
import {
    VictoryPie,
    VictoryBar,
    VictoryTheme,
    VictoryAxis,
    VictoryChart,
    VictoryStack,
    VictoryVoronoiContainer,
    VictoryTooltip,
    VictoryContainer
} from 'victory';
import _ from "lodash";

//ethnicity => number 
//# customers by date 
/*
{
    ethnicity: ''
    count: ''
}
*/
const ethToCus = () => {
    let temp = _.countBy(Customers, 'ethnicity')
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
            label = "ethnicity"
            x="ethnicity"
            y="count"
            colorScale={["tomato", "orange", "gold", "cyan", "navy" ]}
            labelRadius={({ innerRadius }) => innerRadius + 50 }
            // containerComponent={<VictoryContainer responsive={false}/>}
            cornerRadius={({ datum }) => datum.count}
        />
        );
}

const disToCus = () => {
    let temp = _.countBy(Customers, 'disChannel')
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
        containerComponent={<VictoryVoronoiContainer/>}
        // width={600}
        // height={200}
        >
        <VictoryBar
            data={dataToViz}
            label = "disChannel"
            x="disChannel"
            y="count"
            labelRadius={({ innerRadius }) => innerRadius + 50 }
            labelComponent={<VictoryTooltip/>}
            containerComponent={<VictoryVoronoiContainer/>}
            />
            </VictoryChart>
        );
}

const gendToCus = () => {
    
    let temp = _.countBy(Customers, 'gender')
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
            labels={({ datum }) => `${datum.gender}: ${datum.count}`}
            x="gender"
            y="count"
            colorScale={["tomato", "orange", "gold", "cyan", "navy" ]}
            // labelRadius={({ innerRadius }) => innerRadius + 50 }
            containerComponent={<VictoryContainer responsive={true}/>}
            cornerRadius={({ datum }) => datum.count}
        />
        );
}



const Visualization = () => {
    const [visType, setVisType] = useState(null);

    const { customers } = useSelector((state) => state.customer);
    return(
        <div className="container">
        <h1>
            Visualize your customers type
        </h1>
        <button onClick={()=>setVisType(gendToCus)}>Gender</button>
        <button onClick={()=>setVisType(disToCus)}>Discovered Channel</button>
        <button onClick={()=>setVisType(ethToCus)}>Ethnicity</button>
            {visType}
        </div>
    );
}

export default Visualization;