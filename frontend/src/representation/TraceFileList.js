import React, { Component } from 'react';
import PropTypes from 'prop-types'
import moment from 'moment';
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
    compareEntries(entry_a, entry_b) {
        var a = entry_a[1]
        var b = entry_b[1]

        var ts_a = moment(a.firstPacket).format('YYYY-MM-DD')
        var ts_b = moment(b.firstPacket).format('YYYY-MM-DD')
        var ip_a = a.mostFrequentIp
        var ip_b = b.mostFrequentIp

        if (ts_a < ts_b) {
            return -1;
        }
        if (ts_a > ts_b) {
            return 1;
        }

        // dates are equal, compare ip second
        if (ip_a < ip_b) {
            return -1;
        }
        if (ip_a > ip_b) {
            return 1;
        }
        return 0;
    }
    render() {

        return (
            <div className="List">
                <h3>Trace Files (Noise)</h3>
                <p>Drag and drop trace files into days of the timeline.</p>
                <ul>
                {Object.entries(this.getVisibleTracefiles())
                    .sort(this.compareEntries)
                    .map(entry => (
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
