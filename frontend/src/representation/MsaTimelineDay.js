import React, { Component } from 'react';
import PropTypes from 'prop-types';
import MsaTimelineEntry from './MsaTimelineEntry';

import './MsaTimeline.css';


class MsaTimelineDay extends Component {

    drag(event) {
        event.dataTransfer.setData('dayId', event.target.id);
    }

    drop(event) {
        event.preventDefault();
        let draggedElementId = event.dataTransfer.getData('attackId');
        let draggedElementType = 'attack';
        if (!draggedElementId) {
            draggedElementId = event.dataTransfer.getData('traceId');
            draggedElementType = 'traceFile';
        }
        if (!draggedElementId) return;

        try {
            const {addMsaEntry, id} = this.props;
            const draggedId = parseInt(draggedElementId, 10);
            addMsaEntry(id, draggedId, draggedElementType);

        } catch(err) {
            console.error("Drag n Drop error:", err);
        }

    }

    allowDrop(event) {
        event.preventDefault();
    }

    handleDoubleClick(event) {
        event.preventDefault();
        if (event.target.className.includes('MsaTimelineEntry')) {
            return;
        }
        this.props.dayDoubleClick(this.props.id);
    }

    render() {
        const {day, id, entryDoubleClick, getTraceFilesById} = this.props;
        return (
            <div className='MsaTimelineDay' id={id} draggable={true} onDragStart={this.drag} onDrop={this.drop.bind(this)} onDragOver={this.allowDrop} onDoubleClick={this.handleDoubleClick.bind(this)}>
                <h3>Day {id+1}</h3>
                {day.map((entry, idx) => <MsaTimelineEntry key={idx} entry={entry} onDoubleClick={(entryId) => entryDoubleClick(id, entryId)} getTraceFilesById={getTraceFilesById}/>)}
            </div>
        );
    }
}

MsaTimelineDay.propTypes = {
    day: PropTypes.array.isRequired,
    id: PropTypes.number.isRequired,
    getTraceFilesById: PropTypes.func.isRequired,
    entryDoubleClick: PropTypes.func.isRequired,
    dayDoubleClick: PropTypes.func.isRequired,
    addMsaEntry: PropTypes.func.isRequired,
}

export default MsaTimelineDay;