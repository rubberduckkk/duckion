set -e

cd `dirname "$0"`/../ # change to root of credit-card project

rm -rf output # clean previous builds
mkdir -p output/bin output/conf
cp scripts/run_svc.sh output/bootstrap.sh
cp -r config/ output/

cd cmd/credit-server/
go build -v
cd -
mv cmd/credit-server/credit-server output/bin/credit-server
