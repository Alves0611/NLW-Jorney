from flask import jsonify, Blueprint

trips_routes_bp = Blueprint("trip_routes", __name__)

from src.controllers.trip_creator import TripCreator
from src.models.repositories.trips_repository import TripsRepository
from src.models.repositories.emails_to_invite_repository import EmailsToInviteRepository
from src.models.settings.db_connection_handler import db_connection_handler


@trips_routes_bp.route("/trips", methods=["POST"])
def create_trip():
    return jsonify({"hello": "world"}), 200
