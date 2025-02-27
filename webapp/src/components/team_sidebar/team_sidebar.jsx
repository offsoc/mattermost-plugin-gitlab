// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import PropTypes from 'prop-types';

import SidebarButtons from '../sidebar_buttons';

export default class TeamSidebar extends React.PureComponent {
    static propTypes = {
        show: PropTypes.bool.isRequired,
        theme: PropTypes.object.isRequired,
    };

    render() {
        if (!this.props.show) {
            return null;
        }

        return (
            <SidebarButtons
                theme={this.props.theme}
                isTeamSidebar={true}
            />
        );
    }
}
