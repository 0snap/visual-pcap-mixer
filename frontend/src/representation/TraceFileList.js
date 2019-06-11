import React, { Component } from 'react';
import PropTypes from 'prop-types'
import TraceFileListItem from './TraceFileListItem'
import './List.css'

class TraceFileList extends Component {

    getVisibleTracefiles() {
        const traceFiles = {...this.props.traceFiles};
        Object.keys(this.props.traceFiles).forEach(key => {
            if (traceFiles[key].attackTrace) {
                delete traceFiles[key];
            }
        });
        return traceFiles;
    }
    render() {

        return (
            <div className="List">
                <h3>Trace Files (Noise)</h3>
                <p>Drag and drop trace files into days of the timeline.</p>
                <ul>
                {Object.entries(this.getVisibleTracefiles()).map(entry => (
                    <TraceFileListItem key={entry[0]} {...entry[1]} />
                ))}
                </ul>
            </div>
        );
    }
}

TraceFileList.propTypes = {
    traceFiles: PropTypes.object.isRequired,
}

export default TraceFileList;