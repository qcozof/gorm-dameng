#!/bin/bash
git filter-branch --env-filter '
if [ "$GIT_AUTHOR_NAME" = "q.wu" ]
then
export GIT_AUTHOR_NAME="q"
export GIT_AUTHOR_EMAIL="q@cozof.com"
fi
' ref..HEAD

git filter-branch --env-filter '
if [ "$GIT_COMMITTER_NAME" = "q.wu" ]
then
export GIT_COMMITTER_NAME="q"
export GIT_COMMITTER_EMAIL="q@cozof.com"
fi
' ref..HEAD

