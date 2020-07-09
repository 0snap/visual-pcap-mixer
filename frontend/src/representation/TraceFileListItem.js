import React, { Component } from 'react'
import PropTypes from 'prop-types'
import moment from 'moment';
import constants from './constants';
import classNames from 'classnames';

import styles from './List.css';


class TraceFileListItem extends Component {

    drag(event) {
        event.dataTransfer.setData('traceId', event.target.id);
    }

    render() {
        const cx = classNames.bind(styles);
        const classes = cx({
            disabled: this.props.added,
            enabled: !this.props.added,
        });
        const {firstPacket, lastPacket, id, packets, packetsPerSecond, mostFrequentIp} = this.props;
        return (
            <li className={classes} id={id} draggable={true} onDragStart={this.drag}>
                <span>{mostFrequentIp}</span>
                <span>{moment(firstPacket).format(constants.TIME_FORMAT)}&nbsp;&nbsp;&nbsp;{'→'}&nbsp;&nbsp;&nbsp;{moment(lastPacket).format(constants.TIME_FORMAT)}</span>
                <br/>
                <span>Total {packets}&nbsp;pkt</span>
                <span>{packetsPerSecond}&nbsp;pkt/s</span>
            </li>
        );
    }
}

TraceFileListItem.propTypes = {
    path: PropTypes.string.isRequired,
    packets: PropTypes.number.isRequired,
    firstPacket: PropTypes.string.isRequired,
    lastPacket: PropTypes.string.isRequired,
    packetsPerSecond: PropTypes.number.isRequired,
    bytesPerSecond: PropTypes.number.isRequired,
    id: PropTypes.number.isRequired,
    added: PropTypes.bool.isRequired,
}

export default TraceFileListItem;
