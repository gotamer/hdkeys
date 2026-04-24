# file must be sourced not run
# . setup.sh

if [ ! -f "$PWD/setup.sh" ]
then
    echo '!!! You must be in the same folder as your setup.sh file !!!'
    echo "You are in $PWD "
    return
fi

export HDKEYS_ROOT_FOLDER="$PWD"

export HDKEYS_MNEMONIC='tree matrix federal video roof sniff great wheel valve rib daughter public'
export HDKEYS_PASSWORD="somePass@word"

# For hdkeys_wordlist
export HDKEYS_TARGET_KEY="3CN3xb1NFmUFZGQMMvS68jDcoi19RqC19S"
export HDKEYS_PASSWORD_MAX_LENGTH_WORDS='4'
export HDKEYS_PASSWORD_WORDS_SEPERATOR='@'

alias cd_setup=$PWD

