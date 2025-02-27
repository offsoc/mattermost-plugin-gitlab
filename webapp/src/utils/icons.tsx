// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';

type IconProps = {
    fill: string;
    width?: number;
    height?: number;
}

export const GitLabIssuesIcon = ({fill, width = 13, height = 13}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M3 2.5h6a.5.5 0 0 1 .5.5v10a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5ZM1 3a2 2 0 0 1 2-2h6a2 2 0 0 1 1.97 1.658l2.913 1.516a1.75 1.75 0 0 1 .744 2.36l-3.878 7.45a.753.753 0 0 1-.098.145c-.36.526-.965.871-1.651.871H3a2 2 0 0 1-2-2V3Zm10 7.254 2.297-4.413a.25.25 0 0 0-.106-.337L11 4.364v5.89Z'
            fill={fill}
        />
    </svg>
);

export const GitLabMergeRequestIcon = ({fill, width = 13, height = 13}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M10.34 1.22a.75.75 0 0 0-1.06 0L7.53 2.97 7 3.5l.53.53 1.75 1.75a.75.75 0 1 0 1.06-1.06l-.47-.47h.63c.69 0 1.25.56 1.25 1.25v4.614a2.501 2.501 0 1 0 1.5 0V5.5a2.75 2.75 0 0 0-2.75-2.75h-.63l.47-.47a.75.75 0 0 0 0-1.06ZM13.5 12.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Zm-9 0a1 1 0 1 1-2 0 1 1 0 0 1 2 0Zm1.5 0a2.5 2.5 0 1 1-3.25-2.386V5.886a2.501 2.501 0 1 1 1.5 0v4.228A2.501 2.501 0 0 1 6 12.5Zm-1.5-9a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z'
            fill={fill}
        />
    </svg>
);

export const GitLabReviewsIcon = ({fill, width = 13, height = 13}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M9 2.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Zm1.45-.5a2.5 2.5 0 0 0-4.9 0H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1h-2.55ZM8 5H5.5V3.5h-2v11h9v-11h-2V5H8ZM5 7.75A.75.75 0 0 1 5.75 7h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 7.75Zm.75 1.75a.75.75 0 0 0 0 1.5h4.5a.75.75 0 0 0 0-1.5h-4.5Z'
            fill={fill}
        />
    </svg>
);

export const GitLabTodosIcon = ({fill, width = 13, height = 13}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M3 13.5a.5.5 0 0 1-.5-.5V3a.5.5 0 0 1 .5-.5h9.25a.75.75 0 0 0 0-1.5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V9.75a.75.75 0 0 0-1.5 0V13a.5.5 0 0 1-.5.5H3Zm12.78-8.82a.75.75 0 0 0-1.06-1.06L9.162 9.177 7.289 7.241a.75.75 0 1 0-1.078 1.043l2.403 2.484a.75.75 0 0 0 1.07.01L15.78 4.68Z'
            fill={fill}
        />
    </svg>
);

export const GitLabMergeRequestClosedIcon = ({fill, width = 16, height = 16}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M1.22 1.22a.75.75 0 0 1 1.06 0L3.5 2.44l1.22-1.22a.75.75 0 0 1 1.06 1.06L4.56 3.5l1.22 1.22a.75.75 0 0 1-1.06 1.06L3.5 4.56 2.28 5.78a.75.75 0 0 1-1.06-1.06L2.44 3.5 1.22 2.28a.75.75 0 0 1 0-1.06ZM7.5 3.5a.75.75 0 0 1 .75-.75h2.25a2.75 2.75 0 0 1 2.75 2.75v4.614a2.501 2.501 0 1 1-1.5 0V5.5c0-.69-.56-1.25-1.25-1.25H8.25a.75.75 0 0 1-.75-.75Zm5 10a1 1 0 1 0 0-2 1 1 0 0 0 0 2Zm-8-1a1 1 0 1 1-2 0 1 1 0 0 1 2 0Zm1.5 0a2.5 2.5 0 1 1-3.25-2.386V7.75a.75.75 0 0 1 1.5 0v2.364A2.501 2.501 0 0 1 6 12.5Z'
            fill={fill}
        />
    </svg>
);

export const GitLabMergedIcon = ({fill, width = 16, height = 16}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M5.5 3.5a1 1 0 1 1-2 0 1 1 0 0 1 2 0Zm-.044 2.31a2.5 2.5 0 1 0-1.706.076v4.228a2.501 2.501 0 1 0 1.5 0V8.373a5.735 5.735 0 0 0 3.86 1.864 2.501 2.501 0 1 0 .01-1.504 4.254 4.254 0 0 1-3.664-2.922ZM11.5 10.5a1 1 0 1 0 0-2 1 1 0 0 0 0 2Zm-6 2a1 1 0 1 1-2 0 1 1 0 0 1 2 0Z'
            fill={fill}
        />
    </svg>
);

export const GitLabIssueClosedIcon = ({fill, width = 16, height = 16}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M14.5 8a6.5 6.5 0 1 1-13 0 6.5 6.5 0 0 1 13 0ZM16 8A8 8 0 1 1 0 8a8 8 0 0 1 16 0ZM3.75 7.25a.75.75 0 0 0 0 1.5h8.5a.75.75 0 0 0 0-1.5h-8.5Z'
            fill={fill}
        />
    </svg>
);

export const GitLabIssueOpenIcon = ({fill, width = 16, height = 16}: IconProps) => (
    <svg
        width={width}
        height={height}
        viewBox='0 0 16 16'
        xmlns='http://www.w3.org/2000/svg'
    >
        <path
            fillRule='evenodd'
            clipRule='evenodd'
            d='M8 14.5a6.5 6.5 0 1 0 0-13 6.5 6.5 0 0 0 0 13ZM8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16Z'
            fill={fill}
        />
    </svg>
);
