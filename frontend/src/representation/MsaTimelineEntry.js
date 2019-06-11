import React, { Component } from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import moment from 'moment';
import constants from './constants';

import styles from './MsaTimeline.css';


class MsaTimelineEntry extends Component {

    getTooltipText() {
        const {type} = this.props.entry;
        if (type === 'attack')
            return this.props.entry.traces.map(traceId => this.getTracefileTooltipText(this.props.getTraceFilesById(traceId)));

        return this.getTracefileTooltipText(this.props.entry)
    }

    getTracefileTooltipText(tf) {
        if (!tf) {
            return 'undefined traceFile';
        }
        const {id, firstPacket, lastPacket, packets, packetsPerSecond, mostFrequentIp} = tf;
        return (
            <div className='TraceFileDescription' key={tf.id}>
                <span>Tracefile {id}</span>
                <span>Main IP: {mostFrequentIp}</span>
                <span>{moment(firstPacket).format(constants.TIME_FORMAT)} → {moment(lastPacket).format(constants.TIME_FORMAT)}</span>
                <span>{packets} packets</span>
                <span>{packetsPerSecond} pkt/s</span>
            </div>
        );
    }

    render() {
        const {type, id} = this.props.entry;
        const cx = classNames.bind(styles);
        const classes = cx({
            attack: type === 'attack',
            MsaTimelineEntry: true,
        });
        const displayText = type === 'attack' ? this.props.entry.name : this.props.entry.mostFrequentIp;
        return (
            <div className={classes} onDoubleClick={() => this.props.onDoubleClick(id)}>
                {displayText}
                <div className='tooltip'>
                    {this.getTooltipText()}
                </div>
            </div>
        );
    }
}

MsaTimelineEntry.propTypes = {
    entry: PropTypes.object.isRequired,
    getTraceFilesById: PropTypes.func.isRequired,
    onDoubleClick: PropTypes.func.isRequired,
}

export default MsaTimelineEntry;