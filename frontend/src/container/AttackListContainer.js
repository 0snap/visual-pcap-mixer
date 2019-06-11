import { connect } from 'react-redux'
import AttackList from '../representation/AttackList';

function getVisibleAttacks(attacks) {
    // TODO: fixme
    return attacks;
}

const mapStateToProps = state => {
    return {
        attacks: getVisibleAttacks(state.attacks)
    }
}

const mapDispatchToProps = {};

const AttackListContainer = connect(
    mapStateToProps,
    mapDispatchToProps
)(AttackList)

export default AttackListContainer